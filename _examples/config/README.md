#### Config POC

The goal here was to write golang to re-create operations found in krew's `Exec()` which calls `os.Commands()` as a git-wrapper.

Config fetches specifically the `remote.origin.url`, but can be used for other things.
