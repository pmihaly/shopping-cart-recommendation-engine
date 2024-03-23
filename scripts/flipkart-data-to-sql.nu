def main [] {
	open flipkart_com-ecommerce_sample.csv
	| where retail_price != ""
	| where {($in.product_name | str length) <= 120}
	| update product_category_tree {from json | $in.0 | split row " >> "}
	| where {($in.product_category_tree | length) > 3}
	| update image {from json | $in.0}
	| update retail_price {into int}
	| par-each { $"\('($in.product_name | str replace --all `'` `''`)', '($in.description | str replace --all `'` `''`)', '($in.product_category_tree | to json -r | str replace --all `'` `''` | str replace `[` `{` | str replace `]` `}`)', '($in.image)', ($in.retail_price)\)," }
	| prepend "insert into product (name, description, category, image_url, price) values"
	| [...($in | drop), ($in | last | str reverse | str replace ',' ';' | str reverse )]
	| to text
	| save seed/products/01-products-flipkart.sql -f
}
