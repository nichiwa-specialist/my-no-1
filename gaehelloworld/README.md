# GAE

Demonstrate a simple Google App Engine app

Here are my steps to make it work with the GAE SDK.
(Probably not the best ones)

Assuming that go-json-rest is installed using "go get"
and that the GAE SDK is also installed.

Setup:
 * copy this examples/gae/ dir outside of the go-json-rest/ tree
 * cd gae/
 * mkdir -p github.com/ant0ine
 * cp -r $GOPATH/src/github.com/ant0ine/go-json-rest github.com/ant0ine/go-json-rest
 * rm -rf github.com/ant0ine/go-json-rest/examples/
 * path/to/google_appengine/dev_appserver.py .

curl demo:
```
curl -i http://127.0.0.1:8080/message
```
