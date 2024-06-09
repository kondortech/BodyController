'use client'
import React, { useEffect, useState } from "react";
import IngredientCard from "./ingredient_card";
import { ApiListIngredientsResponse, ModelsIngredient } from "@/generated/services/nutrition/api";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import { listIngredients } from "@/services/nutrition/ingredients_api";
import Head from "next/head";

export default function Page() {
	const [ingredientsState, updateIngredients] = useState<ModelsIngredient[]>();

	useEffect(() => {
		listIngredients().then((resp: ApiListIngredientsResponse) => {
			updateIngredients(resp.entities);
		});
	}, []);
	const router = useRouter();

	return (
		<div>
			<Button onClick={() => { router.push('/nutrition/ingredients/create'); console.log("click registered"); }}>
				Create Ingredient
			</Button>
			<div className="min-h-screen bg-gray-100 py-10">
				<Head>
					<title>Ingredients List</title>
				</Head>
				<div className="container mx-auto px-4">
					<h1 className="text-4xl font-bold text-center mb-8">Ingredients List</h1>
					<div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
						{ingredientsState?.map((ingredient: ModelsIngredient) => (
							<IngredientCard key={ingredient.id} ingredient={ingredient} />
						))}
					</div>
				</div>
			</div>
		</div>
	);
}
