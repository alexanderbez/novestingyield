# Withholding Vesting Account Yield

## Background

This repo demonstrates how to **extend** the `x/distribution` module of the Cosmos
SDK, the default module responsible for tracking and managing staking yield, without
having to fork the Cosmos SDK.

Specifically, it demonstrates how to extend the `x/distribution` module to
prohibit the withdrawal of staking yield rewards from a delegator's account based
on some condition. In this case, the condition is that the delegator has some vesting
tokens remaining in their account. However, it can be extended beyond this.

## Usage

You'll notice this repo's version of `x/distribution` simply extends or embeds the
original `x/distribution` module from the Cosmos SDK and thus is quite small.

This allows us to override existing functionality while being able to still
utilize it. Specifically, we override the `WithdrawDelegationRewards` method --
the topmost method responsible for distributing or withdrawing rewards to delegators.

To use this in an application, you'd simply replace the `x/distribution` module
keeper with the one provided here or simply copy the example in this repo to your
own source code. All that's needed is to use the extended keeper. Everything else
remains the same!
