import asyncio
import itertools
import os
import random
from math import floor
from uuid import uuid4

from httpx import AsyncClient
from tqdm.asyncio import tqdm_asyncio

profiles = [
    "spongebob",
    "barbie",
    "ford fiesta",
    "earrings",
    *map(
        " ".join,
        itertools.product(
            ["watch", "shoes", "jeans", "shirt", "sweatshirt"],
            ["women", "men", "kids"],
        ),
    ),
    "gardening",
    "pottery",
    "nail",
    "baby",
    "office",
]

SERVICE_URL = os.getenv("SERVICE_URL", "http://127.0.0.1:8090")


async def get_products_for_profile(search_term: str) -> list[str]:
    async with AsyncClient() as client:
        res = await client.get(
            f"{SERVICE_URL}/products/search",
            params=dict(q=search_term, take=500),
        )
        res.raise_for_status()

    return [item["ID"] for item in res.json()["Items"]]


async def checkout_cart(items: list[str], semaphore: asyncio.Semaphore):
    async with semaphore, AsyncClient() as client:
        cart_id = uuid4()

        for item in items:
            res = await client.put(
                f"{SERVICE_URL}/carts/{cart_id}/items/{item}",
            )
            res.raise_for_status()

        res = await client.post(
            f"{SERVICE_URL}/carts/{cart_id}/checkout",
        )
        res.raise_for_status()


def get_cart_number_of_products(len_of_products: int) -> int:
    return random.randint(1, len_of_products)


async def main():
    product_by_profiles = await asyncio.gather(
        *(get_products_for_profile(profile) for profile in profiles)
    )

    carts = itertools.chain(
        *(
            [
                random.sample(cart, min(len(cart), random.randint(1, 5)))
                for _ in range(floor(len(product_by_profiles) / 2))
            ]
            for cart in product_by_profiles
            if len(cart) >= 1
        )
    )

    semaphore = asyncio.Semaphore(500)
    await tqdm_asyncio.gather(*(checkout_cart(cart, semaphore) for cart in carts))


if __name__ == "__main__":
    asyncio.run(main())
