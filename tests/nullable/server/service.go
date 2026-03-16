package server

import "context"

type Service interface {
	VariantA(ctx context.Context, i1 Base) (r1 Base)
	VariantB(ctx context.Context, i1 BCustomType) (r1 BCustomType)
	VariantC(ctx context.Context, i1 BCustomTypes) (r1 BCustomTypes)
	VariantD(ctx context.Context, i1 BCustomTypesMap) (r1 BCustomTypesMap)

	VariantE(ctx context.Context, i1 *Base) (r1 *Base)

	VariantF(ctx context.Context, i1 []*Base) (r1 []*Base)

	VariantG(ctx context.Context, i1 map[string]*Base) (r1 map[string]*Base)

	VariantH(ctx context.Context, i1 Base, i2 *Base, i3 []*Base, i4 map[string]Base) (r1 Base, r2 *Base, r3 []*Base, r4 map[string]Base)
}
