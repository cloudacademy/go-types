all: composite conversion funcs primitives structs

# ======================
# composite

composite: arrays slices maps interfaces

arrays:
	go run ./cmd/composite/collection/arrays.go

slices:
	go run ./cmd/composite/collection/slices.go

maps:
	go run ./cmd/composite/collection/maps.go

interfaces:
	go run ./cmd/composite/interfaces/interfaces.go

# ======================
# conversion

conversion: convert typeconversion typeassertion

convert:
	go run ./cmd/conversion/convert/convert.go

typeconversion:
	go run ./cmd/conversion/typeconversion/typeconversion.go

typeassertion:
	go run ./cmd/conversion/typeassertion/typeassertion.go

# ======================
# funcs

funcs:
	go run ./cmd/funcs/functions.go

# ======================
# primitives

primitives:
	go run ./cmd/primitives/primitives.go

# ======================
# structs

structs:
	go run ./cmd/structs/structs.go