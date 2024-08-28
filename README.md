# CCIP Cross Chain Name Service

This repository contains tests for the **CCIP Cross Chain Name Service** (CCNS) implemented using both Foundry and Hardhat frameworks. The project demonstrates how to register and look up a name on both source and destination chains using Chainlink’s CCIP.

## Project Structure

- **`foundry/`**: Contains Foundry-based tests for the CCIP Cross Chain Name Service.
- **`hardhat/`**: Contains Hardhat-based tests for the CCIP Cross Chain Name Service.

## Requirements

- **Foundry:** Follow the [Foundry installation guide](https://book.getfoundry.sh/getting-started/installation) for setup.
- **Hardhat:** Ensure Node.js and npm are installed.

## Testing

### Foundry

To test using Foundry:

1. **Navigate to the Foundry directory:**

   ```bash
   cd CCIP-crossChain-nameService/foundry
   ```

2. **Install dependencies:**

   ```bash
   forge install
   ```

3. **Run the tests:**

   ```bash
   forge test -vv
   ```

   **Sample Test Output:**

   ```sh
   shiva-sai@ssb:foundry$ forge test -vv
   [⠒] Compiling...
   No files changed, compilation skipped

   Ran 1 test for test/test.t.sol:CCIPTest
   [PASS] testRegisterAndLookup() (gas: 117497)
   Logs:
     Alice address: 0x328809Bc894f92807417D2dAD6b7C998c1aFdac6
     Deployed CCIPLocalSimulator at: 0x5615dEB798BB3E4dFa0139dFa1b3D433Cc23b72f
     Chain Selector: 16015286601757825753
     Source Router Address: 0xfD07C974e33dd1626640bA3a5acF0418FaacCA7a
     Destination Router Address: 0xfD07C974e33dd1626640bA3a5acF0418FaacCA7a
     Deployed Source Register at: 0xF62849F9A0B5Bf2913b396098F7c7019b51A820a
     Deployed Source Lookup at: 0x2e234DAe75C793f67A35089C9d99245E1C58470b
     Deployed Destination Lookup at: 0x5991A2dF15A8F6A256D3Ec51E99254Cd3fb576A9
     Deployed Receiver at: 0xc7183455a4C133Ae270771860664b6B7ec320bB1
     Set CrossChainNameService addresses on source lookup and destination lookup
     Enabled chain on Source Register with chainSelector: 16015286601757825753
     Registered Alice with 'alice.ccns' on source chain
     Source address lookup result: 0x328809Bc894f92807417D2dAD6b7C998c1aFdac6
     Destination address lookup result: 0x328809Bc894f92807417D2dAD6b7C998c1aFdac6

   Suite result: ok. 1 passed; 0 failed; 0 skipped; finished in 35.65ms (3.40ms CPU time)
   ```

### Hardhat

To test using Hardhat:

1. **Navigate to the Hardhat directory:**

   ```bash
   cd CCIP-crossChain-nameService/hardhat
   ```

2. **Install dependencies:**

   ```bash
   npm install
   ```

3. **Run the tests:**

   ```bash
   npx hardhat test --network hardhat
   ```

   **Sample Test Output:**

   ```c++
   shiva-sai@ssb:hardhat$ npx hardhat test --network hardhat

     CCIP Cross Chain Name Service - Register and Lookup Alice
     Alice's address:  0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
     CCIP Local simulator address:  0x5FbDB2315678afecb367f032d93F642f64180aa3
     Source Lookup contract address:  0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
     Source Register contract address:  0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
     Destination Lookup contract address:  0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
     Receiver contract address:  0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
     Registered Alice with 'alice.ccns' on source chain
     Source chain lookup address:  0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
     Destination chain lookup address:  0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
       ✔ Should register alice.ccns on the source chain and verify her address on both chains (40ms)

   1 passing (2s)
   ```
