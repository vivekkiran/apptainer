
if test -z "$SINGULARITY_INIT"; then
    PATH=$PATH:/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin
    PS1="Singularity.$SINGULARITY_CONTAINER> "
    SINGULARITY_INIT=1
    export PATH PS1 SINGULARITY_INIT
fi
