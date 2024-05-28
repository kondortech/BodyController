'use client'
import React, { useEffect } from "react";
import { IngredientCard } from "./ingredient_card";
import styles from './styles.module.css';
import { ApiListIngredientsResponse, ModelsIngredient } from "@/generated/services/nutrition/api";
import { fetchIngredients } from "./fetch";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function Page() {
	// useEffect(() => {
	// 	fetchIngredients().then((resp: ApiListIngredientsResponse) => {
	// 		console.log(resp);
	// 	});
	// }, []);

	const ingredients: ModelsIngredient[] = [
		{
			title: "Ham",
			macrosNormalized: {
				calories: 100,
				proteins: 22,
				carbs: 5,
				fats: 3,
			},
		},
		// {
		// 	title: "White bread",
		// 	macrosNormalized: {
		// 		calories: 360,
		// 		proteins: 5.5,
		// 		carbs: 62.4,
		// 		fats: 6.9,
		// 	},
		// },
		// {
		// 	title: "White bread",
		// 	macrosNormalized: {
		// 		calories: 360,
		// 		proteins: 5.5,
		// 		carbs: 62.4,
		// 		fats: 6.9,
		// 	},
		// },
		// {
		// 	title: "White bread",
		// 	macrosNormalized: {
		// 		calories: 360,
		// 		proteins: 5.5,
		// 		carbs: 62.4,
		// 		fats: 6.9,
		// 	},
		// },
	];

	const router = useRouter();

	return (
		<div>

			<Button onClick={() => { router.push('/nutrition/ingredients/create'); console.log("click registered"); }}>
				Create Ingredient
			</Button>
			<p className={styles.page_title}>Available Ingredients</p>
			<div className={styles.grid_container}>
				{ingredients.map((value: ModelsIngredient) => {
					return (
						<IngredientCard ingredient={value} />
					)
				})}
			</div>
		</div>
	);
}
