COLORS="#4E4E4E4E4E4E:#FFFF6B6B6060:#FAFAB0B03636:#FFFFFFFFB6B6:#56569696EDED:#FFFF7373FDFD:#8E8EE4E47878:#EEEEEEEEEEEE:#4F4F4F4F4F4F:#F9F968686060:#FAFAB0B03636:#FDFDFFFFB8B8:#6B6B9F9FEDED:#FCFC6E6EF9F9:#8E8EE4E47878:#FFFFFFFFFFFF"
FOREGROUND="#F8F8F8F8F2F2"
BACKGROUND="#2B2B3E3E5050"

gconftool-2 -s -t string /apps/guake/style/background/color $BACKGROUND
gconftool-2 -s -t string /apps/guake/style/font/palette $COLORS
gconftool-2 -s -t string /apps/guake/style/font/color $FOREGROUND


