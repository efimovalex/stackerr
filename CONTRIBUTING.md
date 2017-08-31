# How to contribute

This file is intended to help the contributors of stackerr with guidelines when
making changes and necessary steps before submitting the PR.

## Getting Started

* Fork this repository on GitHub if you haven't already.

## Making Changes

* Create a topic branch from where you want to base your work.
  * This is usually the master branch.
  * The preferred way of naming the branch is SmallDescription e.g.
    AddContributingMd
  * Only target feature/release branch if you are certain your fix must be on
    that branch.
* Execute tests before committing. On how to run the tests, see the Test section.
* Make commits of logical units.
* Check for unnecessary whitespace with `git diff --check` before committing.
* Make sure your commit messages are clear and well descriptive.

````
    Add CONTRIBUTING.md file to stackerr

    Adding CONTRIBUTING.md file with appropriate details to make it easy for the
    users to contribute to stackerr
````

## Testing your changes

Unit tests must complete successfully before you submit your changes.

### Unit tests

#### Adding unittests

Appropriate unit tests must be added for any new functions, and necessary
updates be made to existing tests for any updates to the existing functions. 

In case of removing an functions or converting it to a No-Op functions, the
corresponding unit tests must be removed.

#### Executing unit tests

Unit tests can be run by the following command:

```sh
go test -v ./...
```

## Submitting Changes

* Push your changes to a topic branch in your fork of the repository.
* Run the unit tests (locally or through a PR builder).
* Before submitting the PR, you should be making sure that the unit tests and
  the acceptance tests pass locally
* Submit a pull request to the repository in the using the
  following PR template
  * Replace the content between the '{' and '}' with something meaningful.  If
    there is nothing to add, replace it with N/A or None.

  ````
  # Summary
  {Include information about what is being done in this merge and why. This
  should be easy to read.}

  ## Risks:
  {Risks that are associated with this merge.  Things that it could possibly
  take down in production if released.}

  ````

* Get the changes functionally tested by a QE in the local environment.
  * This should be done BEFORE merging the PR
* Update the Jira ticket to mark that the changes are ready to be reviewed.
  * Send a request to team-pi@sendgrid.com for a review.
* Once the changes are reviewed and tested, the PR will be merged to master.

## PR Builder

Make sure the PR builder returns success

## Reviewing the Pull Request

The person who reviews the pull request is equally responsible and accountable
for the changes being made.  In addition to the code review guidelines (to be
documented in the wiki and to be linked here), the reviewer should make sure of
the following:

* The tests pass on their local environment / PR builder returns success.
* Validate the updates to the version & CHANGELOG and (if any) README updates.

## References

* This document has been created by referring to the
  [CONTRIBUTING.md](https://github.com/puppetlabs/puppet/blob/master/CONTRIBUTING.md)
  file of puppetlabs/puppet GitHub project
