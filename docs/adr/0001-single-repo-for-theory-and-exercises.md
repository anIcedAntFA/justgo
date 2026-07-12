# Keep theory and exercises in a single repo

This is a solo learning repo, so we keep the guide (markdown theory) and all
runnable code (chapter exercises, milestone projects) in **one repository** rather
than splitting exercises into git submodules or sibling repos. Submodules exist to
share code across repos or give a component its own release lifecycle — neither
applies here, and they add real friction (detached HEAD, double commits, recursive
clone) with no payoff at this scale. Co-locating theory and code lets chapters
cross-reference their exercises and lets `go test ./...` cover everything at once.

The one deliberate exception: if a milestone project (e.g. `gitm`) later matures
into a tool worth `go install`-ing and publishing on its own, it graduates to its
own repository — it does not become a submodule of this one.
