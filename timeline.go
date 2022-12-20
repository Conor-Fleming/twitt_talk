package main

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/filteredstream"
	"github.com/michimani/gotwi/tweet/filteredstream/types"
	"github.com/michimani/gotwi/tweet/timeline/types"
)

func CreateSearchRules(client *gotwi.Client) {
	p := &types.CreateRulesInput{
		Add: []types.AddingRule{
			{Value: gotwi.String("@AutomatedAndy"), Tag: gotwi.String("Tweets to Andy rule")},
		},
	}

	res, err := filteredstream.CreateRules(context.TODO(), client, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range res.Data {
		fmt.Printf("ID: %s, Val: %s, Tag: %s\n", gotwi.StringValue(v.ID), gotwi.StringValue(v.Value), gotwi.StringValue(v.Tag))
	}
}
