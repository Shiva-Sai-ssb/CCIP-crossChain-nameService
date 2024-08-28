import { ethers } from "hardhat";
import { expect } from "chai";

describe("CCIP Cross Chain Name Service - Register and Lookup Alice", function () {
  let localSimulator: any;
  let config: any;
  let sourceLookup: any;
  let sourceRegister: any;
  let destinationLookup: any;
  let receiver: any;
  let alice: any;

  before(async function () {
    // Get the signers
    [alice] = await ethers.getSigners();
    console.log("Alice's address: ", alice.address);

    // Deploy CCIPLocalSimulator contract
    const LocalSimulatorFactory = await ethers.getContractFactory(
      "CCIPLocalSimulator"
    );
    localSimulator = await LocalSimulatorFactory.deploy();
    console.log("CCIP Local simulator address: ", localSimulator.address);

    // Retrieve Router contract address from configuration()
    config = await localSimulator.configuration();

    // Deploy contracts on the source chain
    const LookupFactory = await ethers.getContractFactory(
      "CrossChainNameServiceLookup"
    );
    sourceLookup = await LookupFactory.deploy();
    console.log("Source Lookup contract address: ", sourceLookup.address);

    const RegisterFactory = await ethers.getContractFactory(
      "CrossChainNameServiceRegister"
    );
    sourceRegister = await RegisterFactory.deploy(
      config.sourceRouter_,
      sourceLookup.address
    );
    console.log("Source Register contract address: ", sourceRegister.address);

    // Deploy contracts on the destination chain
    destinationLookup = await LookupFactory.deploy();
    console.log(
      "Destination Lookup contract address: ",
      destinationLookup.address
    );

    const ReceiverFactory = await ethers.getContractFactory(
      "CrossChainNameServiceReceiver"
    );
    receiver = await ReceiverFactory.deploy(
      config.destinationRouter_,
      destinationLookup.address,
      config.chainSelector_
    );
    console.log("Receiver contract address: ", receiver.address);

    // Set CrossChainNameService address on both chains
    await sourceLookup.setCrossChainNameServiceAddress(sourceRegister.address);
    await destinationLookup.setCrossChainNameServiceAddress(receiver.address);

    // Enable chain on CrossChainNameServiceRegister
    await sourceRegister.enableChain(
      config.chainSelector_,
      receiver.address,
      500_000n
    );
  });

  it("Should register alice.ccns on the source chain and verify her address on both chains", async function () {
    // Register Alice in CrossChainNameServiceRegister
    const aliceConnected = sourceRegister.connect(alice);
    await aliceConnected.register("alice.ccns");
    console.log("Registered Alice with 'alice.ccns' on source chain");

    // Lookup Alice's address on the source and destination chains
    const sourceAddress = await sourceLookup.lookup("alice.ccns");
    const destinationAddress = await destinationLookup.lookup("alice.ccns");

    console.log("Source chain lookup address: ", sourceAddress);
    console.log("Destination chain lookup address: ", destinationAddress);

    // Assertions
    expect(sourceAddress).to.equal(alice.address);
    expect(destinationAddress).to.equal(alice.address);
  });
});
