#### Clean POC

The goal here was to write golang to re-create the operations found in krew's `Exec()` which calls `os.Commands()` as a git-wrapper.

This looks for a local working directory, fetches the origin/master, then resets to that commit, finally cleaning anything else.
