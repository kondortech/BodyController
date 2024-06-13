'use client';

import '@/app/globals.css';
import React, { useEffect, useState } from "react";
import { RecipeCard } from "./recipe_card";
import Head from "next/head";
import { ApiListRecipesResponse, ModelsRecipe } from "@/generated/services/nutrition/api";
import { listRecipes } from "@/services/nutrition/api";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";


export default function Page() {
  const [recipes, updateRecipes] = useState<ModelsRecipe[]>([]);

  useEffect(() => {
    listRecipes().then((resp: ApiListRecipesResponse) => {
      updateRecipes(resp.entities);
    });
  }, []);
  const router = useRouter();

  return (
    <div className="min-h-screen bg-gray-100 py-10">
      <Head>
        <title>Recipes</title>
      </Head>
      <div className="container mx-auto px-4">
        <Button onClick={() => { router.push('/nutrition/recipes/create'); }}>
          Create Recipe
        </Button>
        <h1 className="text-4xl font-bold text-center mb-8">Recipes</h1>
        <div className="space-y-8">
          {recipes.map((recipe) => (
            <RecipeCard key={recipe.id} recipe={recipe} />
          ))}
        </div>
      </div>
    </div>
  );
}
