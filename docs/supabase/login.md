## tealbase-login

Connect the Tealbase CLI to your Tealbase account by logging in with your [personal access token](https://tealbase.com/dashboard/account/tokens).

Your access token is stored securely in [native credentials storage](https://github.com/zalando/go-keyring#dependencies). If native credentials storage is unavailable, it will be written to a plain text file at `~/.tealbase/access-token`.

> If this behavior is not desired, such as in a CI environment, you may skip login by specifying the `TEALBASE_ACCESS_TOKEN` environment variable in other commands.

The Tealbase CLI uses the stored token to access Management APIs for projects, functions, secrets, etc.
