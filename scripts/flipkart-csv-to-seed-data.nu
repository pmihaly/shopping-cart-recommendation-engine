def on-last [closure: closure] {
	[...($in | drop), ($in | last | do $closure )]
}


def main [] {
	let products = open flipkart_com-ecommerce_sample.csv
	| where retail_price != ""
	| where {($in.product_name | str length) <= 120}
	| update product_category_tree {from json | $in.0 | split row " >> "}
	| where {($in.product_category_tree | length) > 3}
	| update image {from json | $in.0}
	| update retail_price {into int}

	let productsDfr = $products | dfr into-df

	$products
	| par-each {{"productId:ID":  $in.uniq_id, ":LABEL": "Product"}}
	| to csv
	| save ./seed/neo4j/00-products-flipkart.csv -f

	$products
	| par-each { $"\('($in.uniq_id)' ,'($in.product_name | str replace --all `'` `''`)', '($in.description | str replace --all `'` `''`)', '($in.product_category_tree | to json -r | str replace --all `'` `''` | str replace `[` `{` | str replace `]` `}`)', '($in.image)', ($in.retail_price)\)," }
	| prepend "insert into product (id, name, description, category, image_url, price) values"
	| on-last {str reverse | str replace ',' ';' | str reverse}
	| to text
	| save ./seed/postgres/01-products-flipkart.sql -f

	let carts = seq 1 (($products | length) * 2)
	| par-each {|cartId|
		$productsDfr
		| dfr sample -n (random int 0..5) -s $cartId
		| dfr into-nu
		| {":START_ID": $cartId, , ":END_ID": $in.uniq_id }
	}

	$carts
	| $in.":START_ID"
	| par-each {{"cartId:ID": $in, ":LABEL": "Cart"}}
	| to csv
	| save ./seed/neo4j/01-carts-nodes.csv -f

	$carts
	| flatten
	| to csv
	| save ./seed/neo4j/02-carts-relationships.csv -f

}
