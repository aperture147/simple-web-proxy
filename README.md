# Simple Web Proxy

A simple web proxy which can proxy GET HTTP request through another server. Response are returned as-is, even the redirect one.

## How to use

> Assume that your domain is `proxy.example.com`

* Check for server response:
```
https://proxy.example.com/ping
```

* Proxy a request:
```
https://proxy.example.com?url=https://your.url
```

## How to deploy

The code is super simple and lightweight, you can basically deploy it anywhere you want. To simplify the management process, I suggest some PaaS services:
    
- [Heroku](https://www.heroku.com/): Their free plan (with auto turn off dyno) _got cancelled_. Now it's `Eco Dyno`, which costs 5$ a month for 1000hrs of run time. The only advantage of this is the large library of free add-ons (which we don't use) and 2TB of Egress bandwidth per month.

- [Fly.io](https://fly.io/): Their _"free"_ plan is pretty complicated. They have Pay-as-you-go plan that charges for what you use (including computing and network resource). But if you are a personal user having bill < 5$, it will be waived. But nothing ensure that they will not charge us for that money. If you keep it down low then hopefully they will waive the bill.

- [Render](https://render.com/): They have a truly free plan, which does not require any billing info, with 100G bandwidth per month. The downside of Render is the duration to warm up their free runtime, which often takes 50+ seconds.