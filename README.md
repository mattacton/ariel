# ariel

Look at this stuff. Isn't it neat? Wouldn't you think my collections' are useable?

ariel is a command line tool to more easily work with Postman collections and
the offline scratchpad (https://blog.postman.com/announcing-new-lightweight-postman-api-client/). This tool is not suggested for use if you are able
to sign into Postman and use the fully featured online version.

## beta
This tool is in beta. It is not feature complete and may have bugs.

### Usage
```bash
# usage: ariel <collection directory> | <collection.json>
ariel test/test-collections
```
Point ariel at a directory containing Postman collection JSON files, or to
a single collection JSON file. ariel will then start a session where you can traverse
the collections and requests.

```text
====================HELP====================
<num> - Choose an item
<b> - Go back
<c> - Choose a new collection
<r> - Refresh
<q> - Quit
<?> - Help
============================================


1: Twitter API v2
2: Rando Collection
3: FlexWeather Endpoints
4: OpenAI

Choose a collection:_

```
Select a collection to view the requests or folders within that collection.

```text
1: (GET) Followers of user ID   2: (GET) Users a user ID is following
3: (POST) Follow a user ID      4: (DELETE) Unfollow a user ID

Choose an item:_

```
Choose a request and a curl command of that request will be copied to your clipboard.

```text
1: (GET) Followers of user ID   2: (GET) Users a user ID is following
3: (POST) Follow a user ID      4: (DELETE) Unfollow a user ID


 0> Copied Followers of user ID to clipboard
Choose an item:_
```

### Installation
Grab the correct binary executable for your platform from the releases page and
add it to your path.

### Session Options
`<num>`  Choose an item from the list displayed

`<b>` - Go back up the stack, i.e. either to the previous folder or to select a new collection

`<c>` - Return back to the choose a collection screen

`<r>` - Refresh the current screen

`<q>` - Quit the session

`<?>` - Help - prints this help

