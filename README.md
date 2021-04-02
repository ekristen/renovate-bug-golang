# Renovate Bot Golang Module Bug

This is to demonstrate a bug in renovate with the AWS SDK v2 module and how go modules work.

[Reference](https://golang.org/ref/mod#vcs-version)

> If a module is defined in a subdirectory within the repository, that is, the module subdirectory portion of the module path is not empty, then each tag name must be prefixed with the module subdirectory, followed by a slash. For example, the module golang.org/x/tools/gopls is defined in the gopls subdirectory of the repository with root path golang.org/x/tools. The version v0.4.0 of that module must have the tag named gopls/v0.4.0 in that repository.
