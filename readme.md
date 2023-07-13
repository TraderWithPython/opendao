# opendao

**opendao** 是一个治理并分发激励的 DAO
ignite v0.26.1

## 开始运行

```
ignite chain serve --home ./build/data -r
```

### 前置

链的原生代币是 open
默认有三个委员会成员，分别是
m1, m2, m3.
还有 a1, b1 两个无权限的普通账户。

### 检查申请人 a1 和受益人 b1 的账户

检查申请人的账户

```
opendaod q bank balances opendao17xlgmrel0fs4n2yf85w7kft68wyvk0ucd5g2gs
```

检查受益人的账户

```
opendaod q bank balances opendao16g56wtvcvr5tc49l7gfc2frqxukeew66sled64
```

### 提出和查询议案

任何人都可以提出议题，这里假设 a1 提出议题，a1 需要质押 1 open。
a1 指定 b1 是受益人。
如果提案在 60 个区块内未被通过就会自动过期，不返还 a1 的质押金。
b1 的地址是 opendao16g56wtvcvr5tc49l7gfc2frqxukeew66sled64

```
opendaod tx od propose-send "the first proposal" "offer b1 Open and USDo to develop" opendao16g56wtvcvr5tc49l7gfc2frqxukeew66sled64 "10open,10000USDo" --from a1
```

查询议案

```
opendaod q od list-proposal
```

检查申请人的账户

```
opendaod q bank balances opendao17xlgmrel0fs4n2yf85w7kft68wyvk0ucd5g2gs
```

### 投票议案

只有 mp 议员才能投票，同意过半就会通过

```
opendaod tx od vote 0 --from m1
opendaod q od list-proposal
```

```
opendaod tx od vote 0 --from m2
opendaod q od list-proposal
```

返还申请人质押的资金

```
opendaod q bank balances opendao17xlgmrel0fs4n2yf85w7kft68wyvk0ucd5g2gs
```

为了展示过期的功能，通过的提案也还留着，到期时才会删除

```
opendaod q od list-proposal
```

### 查询 b1 收到的资金

```
opendaod q bank balances opendao16g56wtvcvr5tc49l7gfc2frqxukeew66sled64
```
