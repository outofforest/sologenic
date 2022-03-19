# Freeze module

## Genesis

* Map of addresses allowed to (un)freeze tokens

## Configuration

* Reference to keeper of `bank` module

## State

* Map of frozen balances for each account

## Invariants

* checking that frozen balance is not negative
* checking that sum of frozen balance of all the accounts is equal to balance hold by module account

## Keeper

Cooperates with the keeper of `bank` module.

### `Freeze(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) error`

* calls `DelegateCoinsFromAccountToModule` of `bank` keeper
* increments frozen balance of the account

### `Unfreeze(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) error`

* checks that frozen balance of the account is high enough
* decrements frozen balance
* calls `UndelegateCoinsFromModuleToAccount`

### `Frozen(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins`

* returns frozen balance for the account

## Querier

Querier allowing to query frozen balance of an account

## Messages

### `MsgBalance`

Used for both `freeze` and `unfreeze` operation

```
message MsgBalance {
  option (cosmos.msg.v1.signer) = "sender";

  option (gogoproto.equal)           = false;
  option (gogoproto.goproto_getters) = false;

  string   sender                          = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string   address                         = 2 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  repeated cosmos.base.v1beta1.Coin amount = 3
      [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"];
}
```

## Services

```
service Msg {
  rpc Freeze(MsgBalance) returns (MsgResponse);
  rpc Unfreeze(MsgBalance) returns (MsgResponse);
}
```

Basic operation:
* decode context and addresses
* call `Freeze` or `Unfreeze` method of keeper
* emit metric and event

## Additional features
* transactions for updating addresses having permissions to (un)freeze balances