install: ~/.local/bin/jot ~/.cache/oh-my-zsh/completions/_jot

~/.local/bin/jot: $(wildcard cmd/*.go)
	go build -o $@ ./cmd

~/.cache/oh-my-zsh/completions/_jot: ~/.local/bin/jot
	jot completion zsh > $@

clean:
	rm -f ~/.local/bin/jot
	rm -f ~/.cache/oh-my-zsh/completions/_jot

