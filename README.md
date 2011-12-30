
=============

Pillars
-------

* No Magic
* No Reflect
* No Unsafe
* As functional as feasible, i.e. defering mutable state until optimization phase


Dependecies
-----------

The dependencies listed are required if you wish to use SansMagic

* [.golang](http://golang.org/doc/install.html) 

Contributing
------------

Want to contribute? Great! 



### Getting Started

$ goinstall sansmagic
$ sans new-proj sans-hello
$ cd sans-hello
$ make run

    command(:rest2html, /re?st(.txt)?/)

Here we're telling GitHub Markup of the existence of a `rest2html`
command which should be used for any file ending in `rest`,
`rst`, `rest.txt` or `rst.txt`. Any regular expression will do.

Finally add your tests. Create a `README.extension` in `test/markups`
along with a `README.extension.html`. As you may imagine, the
`README.extension` should be your known input and the
`README.extension.html` should be the desired output.

Now run the tests: `rake`

If nothing complains, congratulations!


### Classes

If your markup can be translated using a Ruby library, that's
great. Check out Check `lib/github/markups.rb` for some
examples. Let's look at Markdown:

    markup(:markdown, /md|mkdn?|markdown/) do |content|
      Markdown.new(content).to_html
    end

We give the `markup` method three bits of information: the name of the
file to `require`, a regular expression for extensions to match, and a
block to run with unformatted markup which should return HTML.

If you need to monkeypatch a RubyGem or something, check out the
included RDoc example.

Tests should be added in the same manner as described under the
`Commands` section.


Installation
-----------

    goinstall sansmagic

if this doesn't work then you probably have an old go install (that is of course
if my code didn't break it!)

Example
-----


Testing
-------

To run the tests:

    $ gotest

To add tests see the `Commands` section earlier in this
README.


Contributing
------------

1. Fork it.
2. Create a branch (`git checkout -b my_branch`)
3. Commit your changes (`git commit -am "recoded the whole thing"`)
4. Push to the branch (`git push origin my_branch`)
5. Create an [Issue][1] with a link to your branch
6. Prepare for eternal fame and fortune.

[1]: http://github.com/drhodes/sansmagic/issues
