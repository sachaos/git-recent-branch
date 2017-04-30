function peco-recent-branch () {
    local branch=$(git recent-branch | peco | cut -d ' ' -f 1 )
    if [ -n "$branch" ]; then
        if [ -n "$LBUFFER" ]; then
            local new_left="${LBUFFER%\ } $branch"
        else
            local new_left="$branch"
        fi
        BUFFER=${new_left}${RBUFFER}
        CURSOR=${#new_left}
    fi
}
zle -N peco-recent-branch
bindkey '^g^r' peco-recent-branch
