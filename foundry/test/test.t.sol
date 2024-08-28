// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {CCIPLocalSimulator, IRouterClient} from "@chainlink/local/src/ccip/CCIPLocalSimulator.sol";
import {CrossChainNameServiceRegister} from "src/CrossChainNameServiceRegister.sol";
import {CrossChainNameServiceLookup} from "src/CrossChainNameServiceLookup.sol";
import {CrossChainNameServiceReceiver} from "src/CrossChainNameServiceReceiver.sol";

contract CCIPTest is Test {
    CCIPLocalSimulator public localSimulator;
    CrossChainNameServiceRegister public sourceRegister;
    CrossChainNameServiceLookup public sourceLookup;
    CrossChainNameServiceLookup public destinationLookup;
    CrossChainNameServiceReceiver public receiver;

    address public alice;

    function setUp() public {
        // Set up addresses
        alice = makeAddr("alice");
        console.log("Alice address:", alice);

        // Deploy CCIPLocalSimulator
        localSimulator = new CCIPLocalSimulator();
        console.log("Deployed CCIPLocalSimulator at:", address(localSimulator));

        // Retrieve configuration details
        (uint64 chainSelector, IRouterClient sourceRouter, IRouterClient destinationRouter,,,,) =
            localSimulator.configuration();
        console.log("Chain Selector:", chainSelector);
        console.log("Source Router Address:", address(sourceRouter));
        console.log("Destination Router Address:", address(destinationRouter));

        // Deploy contracts on the source chain
        sourceLookup = new CrossChainNameServiceLookup();
        sourceRegister = new CrossChainNameServiceRegister(address(sourceRouter), address(sourceLookup));
        console.log("Deployed Source Register at:", address(sourceRegister));
        console.log("Deployed Source Lookup at:", address(sourceLookup));

        // Deploy contracts on the destination chain
        destinationLookup = new CrossChainNameServiceLookup();
        receiver =
            new CrossChainNameServiceReceiver(address(destinationRouter), address(destinationLookup), chainSelector);
        console.log("Deployed Destination Lookup at:", address(destinationLookup));
        console.log("Deployed Receiver at:", address(receiver));

        // Set CrossChainNameService addresses
        sourceLookup.setCrossChainNameServiceAddress(address(sourceRegister));
        destinationLookup.setCrossChainNameServiceAddress(address(receiver));

        console.log("Set CrossChainNameService addresses on source lookup and destination lookup");

        // Enable chain on CrossChainNameServiceRegister
        sourceRegister.enableChain(chainSelector, address(receiver), 500_000);
        console.log("Enabled chain on Source Register with chainSelector:", chainSelector);
    }

    function testRegisterAndLookup() public {
        // Register Alice on the source chain
        vm.startPrank(alice);
        sourceRegister.register("alice.ccns");
        vm.stopPrank();
        console.log("Registered Alice with 'alice.ccns' on source chain");

        // Lookup Alice's address on both chains
        address sourceAddress = sourceLookup.lookup("alice.ccns");
        address destinationAddress = destinationLookup.lookup("alice.ccns");

        console.log("Source address lookup result:", sourceAddress);
        console.log("Destination address lookup result:", destinationAddress);

        // Assertions to verify the addresses
        assertEq(sourceAddress, alice);
        assertEq(destinationAddress, alice);
    }
}
