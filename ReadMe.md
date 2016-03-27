IPFSproxy
=========

(This is a buggy proof of concept. Do not use for production!)

A transparent http proxy that uses [IPFS](https://ipfs.io/) as a cdn. This will allow big mesh networks like [freifunk](https://freifunk.net/) with multiple internet gateways to only download content ones. Static content(images, libraries etc) can be loaded from the mesh without the content provider changing anything.

The main issue that needs to be solved is how do you know the hash of a random link on the internet like cdn.exsample.com/image.jpg. I'm proposing to use trackers much like DNS works today. And that is what has been implemented here.

## Contact
You can find me on IRC #ipfs nickname geir_
