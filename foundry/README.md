# CCIP Cross Chain Name Service - Foundry Test

This repository contains a test suite written in Solidity using the Foundry framework. The test focuses on the **CCIP Cross Chain Name Service** (CCNS) implemented using Chainlink's CCIPLocalSimulator. The test simulates registering a name on a source blockchain and verifying that the name resolves to the correct address on both the source and destination blockchains.

## Overview

### Test Contracts

- **CCIPLocalSimulator:** A local simulator provided by Chainlink to test Cross-Chain Interactions.
- **CrossChainNameServiceRegister:** Handles the registration of names on the source chain.
- **CrossChainNameServiceLookup:** Provides lookup functionality to resolve registered names to addresses.
- **CrossChainNameServiceReceiver:** Receives cross-chain messages and handles name resolution on the destination chain.

### Test Scenario

1. **Setup Phase:**

   - Deploy the `CCIPLocalSimulator` contract.
   - Deploy `CrossChainNameServiceRegister` and `CrossChainNameServiceLookup` on the source chain.
   - Deploy `CrossChainNameServiceReceiver` and `CrossChainNameServiceLookup` on the destination chain.
   - Configure the contracts to enable cross-chain communication.

2. **Testing Phase:**
   - **Register a Name:** Register the name `"alice.ccns"` to the address `alice` on the source chain.
   - **Lookup Name:** Retrieve the registered address for `"alice.ccns"` on both the source and destination chains.
   - **Assertions:** Verify that the address retrieved from both chains matches `alice`'s address.

### Key Functions

- **setUp:** Initializes the test environment by deploying contracts and setting up cross-chain communication.
- **testRegisterAndLookup:** Tests the registration and lookup of a name across the source and destination chains.

## How to Run the Tests

1. **Install Foundry:**
   Follow the [Foundry installation guide](https://book.getfoundry.sh/getting-started/installation) to set up Foundry on your system.

2. **Clone the Repository:**

   ```bash
   git clone https://github.com/Shiva-Sai-ssb/CCIP-crossChain-nameService.git
   cd CCIP-crossChain-nameService && cd foundry
   ```

3. **Install Dependencies:**

   ```bash
   forge install
   ```

4. **Run the Tests:**

   ```bash
   forge test -vv
   ```

   This will execute the test file and display the output, including the logs generated by `console.log` statements.

## Test Logs

The test output includes the following key logs:

- Alice's address on the source chain.
- Deployment addresses of the CCIPLocalSimulator, CrossChainNameServiceRegister, and CrossChainNameServiceLookup contracts.
- Chain selector and router addresses.
- Confirmation of successful registration of `"alice.ccns"`.
- Results of the name lookup on both source and destination chains.

## Test Output

The following is the test output from running the Foundry test:

```bash
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

Ran 1 test suite in 160.25ms (35.65ms CPU time): 1 tests passed, 0 failed, 0 skipped (1 total tests)
```

This detailed output shows the steps taken during the test execution, including the deployment of contracts, configuration of cross-chain communication, and the successful registration and lookup of the `alice.ccns` name on both the source and destination chains.
