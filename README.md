# cli

ClI tool to fetch the translation on the developer machine from [Libraria](https://app.libraria.io/)

The latest builds can be downloaded from [Releases](https://github.com/libraria-app/cli/releases)

This app is a Libraria CLI client that help you to manage your translations just from your terminal.
It allows developer to export term translation from the needed project.

Available Commands:
`export`     Export translation to the file

Export command creates a new or updates exising file by the provided path with the exported project terms translations.

This commands requires to provide:

- the project identificator from Libraria
- the language code from ISO 639 list
- the filename where to export with `{lang}` placeholder, i.e. `mydir/{lang}/file.json`. The file extension should be
  one of: `.json`, `.xml`, `.strings`
- the user API key from [Libraria user profile]()

### To make your own build:

1. Install Go **1.20**.
2. Clone the current repository
3. Run from the home project directory
   ```shell
    CGO_ENABLED=0 go build ./cmd/libraria
    ```

This command will create a binary file for your OS by default.