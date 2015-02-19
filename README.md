# abbrvt
Abbreviates a word by removing vowels and a sentence by using its words' capital letters

To use this library, import as follows:  
```golang
import "github.com/callmekatootie/abbrvt"
```

## Examples
To abbreviate a word,  
```golang
abbrvt.Get("Germany") // Grmny
```

A single word is abbreviated by removing all vowels from it.

To abbreviate a sentence,
```golang
abbrvt.Get("National Aeronautics and Space Administration") // NASA
```

A sentence (multiple words) is abbreviated by taking the first letter of each word, only if the letter
is a capital letter. Any word which starts with a small letter is ignored.

## Points to note
* If a sentence does not contain any word that begins with a capital letter, you get an error
* If a sentence has less than 3 words which begin with a capital letter, you get an error

Inspired / ported from this ruby gem - https://github.com/kachick/abb

## License
MIT
