package srvgrpc_test

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	// ts := srvgrpc.NewServer(context.TODO())
	// res1, err := ts.CardPoints(context.TODO(), &pb.CardPointsRequest{CardNumber: 1})
	// if err != nil {
	// 	t.Logf("err type %t", err)
	// 	t.Fatal(err)
	// }
	// t.Log(res1)

	// res2, err := ts.PointCount(context.TODO(), &pb.PointCountRequest{CardNumber: []uint32{1, 8, 10, 4}})
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(res2)

	// res3, err := ts.CardCompare(context.TODO(), &pb.CardCompareRequest{FirstCard: &pb.ItalianCard{Number: 1, Seed: pb.Seed_CUP}, SecondCard: &pb.ItalianCard{Number: 3, Seed: pb.Seed_COIN}, Briscola: pb.Seed_COIN})
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(res3)
}
