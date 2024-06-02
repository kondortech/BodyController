'use client'
import React, { useEffect, useState } from "react";
import { IngredientCard } from "./ingredient_card";
import styles from './styles.module.css';
import { ApiListIngredientsResponse, ModelsIngredient } from "@/generated/services/nutrition/api";
import { fetchIngredients } from "./fetch";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function Page() {
	const [ingredientsState, updateIngredients] = useState<ModelsIngredient[]>();

	useEffect(() => {
		fetchIngredients().then((resp: ApiListIngredientsResponse) => {
			console.log(resp);
			updateIngredients(resp.entities);
		});
	}, []);
	const router = useRouter();

	return (
		<div>

			<Button onClick={() => { router.push('/nutrition/ingredients/create'); console.log("click registered"); }}>
				Create Ingredient
			</Button>
			<p className={styles.page_title}>Available Ingredients</p>
			<div className={styles.grid_container}>
				{ingredientsState?.map((value: ModelsIngredient) => {
					return (
						<IngredientCard ingredient={value} />
					)
				})}
			</div>
		</div>
	);
}
