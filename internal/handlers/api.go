package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Api(ctx echo.Context) error {
	banner := `
kkkkkkkOOO000Ol'...o000000000000OOOOkkkkkkkkkkkOOOOkkkkkkkOOOkkkOOOOOOOOkkOOOO
kkkkkkkOOOO00k:.  .lO00O000000OOOOOOkkkkkkkkkkkOOOOkkkkkkkkOOkkkkkkOOOkkkkkkOO
kkkkOOOOOOOOOOkdc'';:::ccc:;;;cdxxxkkkkkkkkkxxkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkO
kkkkOOOOOOOOOkOOOkxdoolc,.. ..,ccccc::::::::,,,:lxOkkkkkkkkkkkkkkkkkkkkkkkkkkk
kkkOOO00OOOOOkOOOOOOxo:;....:okOOOOOkxxxdoo;...':ooooodxxxdxkkkkkxxxxxxkkkkkkk
kkOOOOOOOOOOkkOOOOdc,.....;dxolc:cloxOOOOOx:'.;xOOkxddool,.'cdxkxxxxxxxxxkkkkk
OOOOOOOOOOOOkkOOxc'......,okc.......,okOOx;',';dOOkkkkkkkdlldddxxxxxxxxxxkkkkk
OOOOOOOOOkkkkkkx:........:xo.........,x0k:..'',cxkkkkkkkxkkkkkxxxxxxxxxxxxkkkk
OOOOOOOkkkkkkkko'.......'lx:. ... ...'odc....',,lxkxxxxxxxxxxxxxxxxxxxxxxxkkkk
OO00OOOkkkkkkxxl'.......'ld,        .,:,.....'',;lxxxxxxxxxxxxxxxxxxxxxxxxkkkk
OO00OOOkkkkkkxxc'........,;.    ..   ........'',,;oxddxxxxxxxxxxxxxxxxxxxkkkOO
OOOOOOOkkkkkxkd;........... ..        . ......'''':ddddddddddddddddxxxxxxkkOOO
OOOOOOOkkkkxxd;.............           ........''.;odddddddddddddddxxxxxkkkOOO
OOOOkkkkkkkxo;...............          ...........'ldddddddddddddddxxxxkkkkkkk
kkkkkkkkkkxo,................           ..........'lxdddddddddddddxxxxxxkkkkkk
kkkkkkkkkxl'............. ... ......    ..........'lxdddddddddddddxxxxxxxxxkkk
kkxxxxkkxl'.................  ...        ... ......cdddddddddddddddxxxxxxxxxkk
`
	return ctx.String(http.StatusOK, banner)
}
