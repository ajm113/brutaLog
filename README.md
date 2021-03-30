# BrutaLog
---

*Phishing site attack software that spams random logins to targeted website's form*

![GitHub Logo](/preview.gif)

**NOTE: Author nor maintainers are NOT liable for users taking advantage of brutaLog against innocent websites or software!**

If you feel someone is using BrutaLog against your website maliciously. Please contact your system admin or make sure your services are setup to IP block large requests that failed to auth or web services that can properly filter legitamit logins perhaps using frontend methods to ensure proper tokens are passed when a login attempt occures.

## How does BrutaLog work?

Sends random logins via rainbow table of emails and or passwords to targeted form submit URL that acts like a browser POST request that should in-theroy bloat the target phishing website database with useless logins that will also increase the phishing website expenses if the scammer doesn't have proper procations against these type of attacks.

## How to Build

Make sure you have the latest Go 1.16+ installed.

`$ go build .`

## Running example:

Sends a single request to target website:

`$ ./brutaLog -u=https://targetwebsite.com/login.php`

Sends attack on website when verified that the used/other command works.

`& ./brutaLog -u=https://targetwebsite.com/login.php -w=100 -c=0`


## Todos:

* ~~Create handle for manager that keeps tabs on requests sent, error counts, verbose mode, etc...~~
* Improved delay controls if user wants the delay to be stupped instead of random between 0 to x.
* Add random email generator.
* Add keep alive headers and or any missing headers the HTTP client should have.
* Add docker container for easier swarm control.
* Add Wiki!
