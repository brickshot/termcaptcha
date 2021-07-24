# termcaptcha

#### v0.1.0 Alpha

CAPTCHA for the terminal. Generates a pile of VT100 compatible control codes to render strings of characters which are (kinda) hard for a computer to decipher but (kinda) easy for a human.

**Note**: There are many ways these CAPTCHAs could be easily cracked. Don't bet the farm on it.

### What problem is this solving?

You have a public API which is accessible via a CLI. You don't want to require a signup or OAuth or really any kind of access control but you also don't want to be spammed to death by bots.

Requiring the user to complete a CAPTCHA can help solve this problem. Currently almost all CAPTCHAs are web based. This CAPTCHA will run in the terminal.

No CAPTCHA is foolproof (this one in particular!). Ultimately a malicious actor could farm out solving the CAPTCHA to actual humans for money. However it can raise the bar for how much effort is required to do annoying things to your service in an automated way.  

### Usage

1. On a host with a publicly accessible IP address (and prefereably a public hostname), run the service executable. It listens on port 8000 by default.
1. On the client side before calling your API, hit the `/get/{uuid}` route on the CAPTCHA server with a UUID of your choosing.
1. Display the result to the user in a VT100 compatible terminal. The user's terminal can display results so fast that the user will never have time to view it being rendered. You can slow down the display using the following `curl` command:

    ```
      curl --no-buffer --limit-rate 700 localhost:8000/get/3bf9de29-02c4-4fcc-9bac-daa21472c814
    ```
   
   Alternatively you could get the response and display it slowly using something like [slowcat](https://github.com/nikushi/slowcat]slowcat).

1. Ask the user what word they saw.
1. Submit the UUID and the word to `/check/{UUID}?word={word}`
1. If the response is "OK" then the user was correct.
1. If the user was correct then you can pass the UUID into your API along with your protected request.
1. On the backend - your service makes a request to `/verify/{UUID}` to confirm that the user passed the CAPTCHA identified by the UUID. If you get "OK" back then go ahead and fulfill the request.

### Notes

Do not re-use UUIDs. They will only work once. Generate a new one each time you request a CAPTCHA.

If the user fails the CAPTCHA, start over. Don't ask them to solve the same one again.

You only have 1 "check" and one "verify" per UUID. If the check passes then the backend of your service can verify it. After that an error `Too many checks` will be returned.

#### Examples

```bash
curl --no-buffer --limit-rate 1000 localhost:8000/get
```

The word `LEARNING` is displayed. Note it is not just sent as a continuous string but the letters are drawn randomly to the terminal using VT100 escape codes.

### Does this stop someone DOS-ing my service?

No.

