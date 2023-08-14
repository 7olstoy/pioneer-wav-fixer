# pioneer-wav-fixer

## The problem
Many files from bandcamp has the "unsupported file type error" issue with old Pioneer XDJ/CDJ decks.

This is because of WAV_EXT flag into these files.

## The solution
We need to rewrite 20th and 21th bytes of these wav-files to ```01 00``` and the Pioneers will start to understand the files.

## Thanks
Big up to pioneer [forum](https://forums.pioneerdj.com/hc/en-us/community/posts/360043048651-E-8305-unsupported-file-type-error) and [this](https://github.com/camm9909/WavPatcher) old python version.
