## Instantly

Mobile money API client built on top of [BROKER](http://developers.smsgh.com/documentations/unity/broker). Yh the name is misleading becuase mobile money transfers have never been that instant. I believe we will get there soon.

## Motivation

I have been building apps that need mobile money integration and I realise I will probably be doing a lot of repetitive work all the time. Instantly hopes to solve that problem for me and for you too. One API for all my mobile money integrations built on top of SMSGH's BROKER


### Documentation

You, yes you, are a developer who is building the NBA (Next Big App) and the app needs to integrate with mobile money payments in Ghana, then use [Instantly](https://github.com/ngenerio/instantly). First register on [Unity](https://unity.smsgh.com), get your `client_id` and `client_secret` from [Unity](https://unity.smsgh.com/account/api-accounts) and go to [Broker](https://apps.smsgh.com/broker/), create a wallet and go to [token](https://apps.smsgh.com/broker/apitokens) to retrieve your wallet `token`.


#### Configuration

Create a config.yaml in the root directory of this project. Specify the configuration as the example below:

```yaml
SERVER_ENV: development # development or production
API_CLIENT_USERNAME: username # The username for basic auth purposes with instantly api server
API_CLIENT_PASSWORD: password # The password for basic auth purposes with instantly api server
UNITY_CLIENT_ID: jlbtfmji # Your unity client ID from SMSGH
UNITY_CLIENT_SECRET: enpmkuke # Your unity client secret from SMSGH
BROKER_TOKEN: 8b05b816-1ca6-4f60-9ca2-966a15960154 # Your broker token from broker
BROKER_SENDER: Instantly # Your app name
BROKER_CALLBACK_URL: https://<host>:/api/v1/payment/callback # Specify host as the url this app has been hosted at
BROKER_BASE_URL: https://api.smsgh.com/usp # Don't touch this
DB_NAME: postgres # Database name: either `postgres` or `mysql`
DB_PATH: user='' dbname=instantly_test sslmode=disable # Your database connection url
MIGRATIONS_DIR: db/db_ # Don't touch this
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


