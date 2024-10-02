# go-git

A very basic version control system

## What go-git can do

- Can take 'sanpshots' of state of project directory
- Ignore certain files based on:
  - Relative path
  - Absolute path
  - Pattern
- Can revert back and forth 'snapshots'
- These 'snapshots' are shareable

## What go-git cannot do

- Branching out into different versions
- Integrate with cloud repository services

### TODO

- [x] Ignore file types
- [x] Ignore directories
- [x] Init command
- [x] Store in hidden directory
- [x] Rraverse dirs
- [ ] Track with timestammps
- [x] Make cwd as base path
- [ ] Make a cli app
- [ ] Jump back and forth snapshots
- [x] Name snapshots
- [ ] Application name
- [x] Sanitize snapshot names
  - [ ] Stricter sanitization
- [ ] Command to create an ignores file
- [ ] Shouldn't paths be relative to make snapshots shareable?
- [ ] How to handle large files?
- [ ] Error handling
- [ ] Custom print like ``[INFO] go-git init success``
  - [ ] What tags to include? INFO, WARNING, ERROR
- [ ] If not init before snapping, init and snap