# Cloudpayments SDK

## Installation

```sh
go get -u github.com/onemgvv/cloudpayments
```

## Usage

```go
sdk := cloudpayments.New("secret-from-lk", "your-public-id")

params := cloudpayments.CryptoPayRequest{
  Amount: 500.3,
  Currency: cloudpayments.RUB,
  IpAddress: "127.0.0.1",
  CardCryptogramPacket: "eyJUeXBlIjoiQ2xvdWRDYXJkIiwi...",
  Name: "Michael",
  PaymentUrl: "https://my-site.com/payment/cb",
  InvoiceId: "invoice12-333",
  CultureName: "ru-RU",
  AccountId: "10",
  Email: "mail@example.com",
}

resp, err := sdk.PayByCryptogram(params)
if err != nil {
  // handle error
}

// use response
```
