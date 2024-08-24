'use client';

import '@/app/globals.css';
import React, { useEffect, useState } from "react";
import { MealCard } from "./meal_card";
import Head from "next/head";
import { ApiListMealsResponse, ApiListRecipesResponse, ModelsMeal, ModelsRecipe } from "@/generated/services/nutrition/api";
import { listMeals, listRecipes } from "@/services/nutrition/api";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";


export default function Page() {
  const [meals, updateMeals] = useState<ModelsMeal[]>([]);

  useEffect(() => {
    listMeals().then((resp: ApiListMealsResponse) => {
      updateMeals(resp.entities);
    });
  }, []);
  const router = useRouter();

  return (
    <div className="min-h-screen bg-gray-100 py-10">
      <Head>
        <title>Meals</title>
      </Head>
      <div className="container mx-auto px-4">
        <Button onClick={() => { router.push('/nutrition/meals/create'); }}>
          Create Meal
        </Button>
        <h1 className="text-4xl font-bold text-center mb-8">Meals</h1>
        <div className="space-y-8">
          {meals.map((meal) => (
            <MealCard key={meal.id} meal={meal} />
          ))}
        </div>
      </div>
    </div>
  );
}
