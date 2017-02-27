## Instantly

Mobile money API client built on top of [BROKER](http://developers.smsgh.com/documentations/unity/broker). Yh the name is misleading becuase mobile money transfers have never been that instant. I believe we will get there soon.

## Motivation

I have been building apps that need mobile money integration and I realise I will probably be doing a lot of repetitive work all the time. Instantly hopes to solve that problem for me and for you too. One API for all my mobile money integrations built on top of SMSGH's BROKER


### Documentation

You, yes you, are a developer who is building the NBA (Next Big App) and the app needs to integrate with mobile money payments in Ghana, then use [Instantly](https://github.com/ngenerio/instantly). First register on [Unity](https://unity.smsgh.com), get your `client_id` and `client_secret` from [Unity](https://unity.smsgh.com/account/api-accounts) and go to [Broker](https://apps.smsgh.com/broker/), create a wallet and go to [token](https://apps.smsgh.com/broker/apitokens) to retrieve your wallet `token`.


#### Configuration

The repo contains a `config.yaml`, we set all our configuration there:

```yaml
SERVER_ENV: development #The server environment (development, production)
API_CLIENT_USERNAME: username # For basic authentication with instantly api
API_CLIENT_PASSWORD: password # For basic authentication with instantly api
UNITY_CLIENT_ID: instantly_nba # client_id from unity
UNITY_CLIENT_SECRET: instantly_nba_secret # client_secret from unity
BROKER_TOKEN: xxxxx-xxxx-xxxx-xxxx-xxxxxxxx # token from Broker
BROKER_SENDER: Instantly # App name
BROKER_CALLBACK_URL: https://instantly.localtunnel.me/api/v1/payment/callback # Callback URL used when the mobile money transaction fails or succeeds
BROKER_BASE_URL: https://api.smsgh.com/usp # Broker URL
POSTGRES_DBNAME: instantly_test # Database name to store transactions data
POSTGRES_URL: user='' dbname=instantly_test sslmode=disable # Postgres Database connection details
MIGRATIONS_DIR: db/migrations # You don't want to touch this. Contains migrations data
```


#### Running

You need `postgres` and `go` installed on your system to get this running:

#### Manual
- Clone the project to your $GOPATH
- Install the dependencies
- Run

```bash
$ git clone https://github.com/ngenerio/instantly
$ glide install
$ ./instantly
```

#### Using GO GET


```bash
$ go get github.com/ngenerio/instantly
```

#### Awesome Instant API Endpoints

This is the data you need to send to `Instantly`:

```json
{
    "phoneNumber": "2332xxxxxxxx",
    "amount": 50.00,
    "mno": "TIGO",
    "name": "Eugene Asiedu",
    "type": "debit"
}
```

##### Parameters
- phoneNumber - the phone number of the individual
- amount - the account to be debited or credited
- mno - the mobile network operator
- name - the name of the individual
- type - either debit or credit


The payment endpoint is http://`host`:`port`/api/v1/payment. Send the `API_CLIENT_USERNAME` and `API_CLIENT_PASSWORD` as your `Basic Authentication` details.


```curl
$ curl -u API_CLIENT_USERNAME:API_CLIENT_PASSWORD -X POST -d '{"phoneNumber":"2332xxxxxxxx","amount":50.00,"mno":"TIGO","name":"Eugene Asiedu","type":"debit"}' http://<host>:<port>/api/v1/payment
```


