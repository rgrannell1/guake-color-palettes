COLORS="#000000000000:#FFFF00000000:#3737DDDD2121:#FEFEE4E40909:#14146060D2D2:#FFFF00005D5D:#0000BBBBBBBB:#BBBBBBBBBBBB:#555555555555:#F4F40D0D1717:#3B3BCFCF1D1D:#ECECC8C80909:#55555555FFFF:#FFFF5555FFFF:#6A6AE3E3F9F9:#FFFFFFFFFFFF"
FOREGROUND="#FFFFFFFFFFFF"
BACKGROUND="#121226263737"

gconftool-2 -s -t string /apps/guake/style/background/color $BACKGROUND
gconftool-2 -s -t string /apps/guake/style/font/palette $COLORS
gconftool-2 -s -t string /apps/guake/style/font/color $FOREGROUND
