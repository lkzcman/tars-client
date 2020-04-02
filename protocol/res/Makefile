all:
	tars2go -add-servant=false -tarsPath github.com/uugtv/tars-client *.tars
	#override file
	sed -i 's|"endpointf"|"github.com/lkzcman/tars-client/protocol/res/endpointf"|g' queryf/QueryF_IF.go
