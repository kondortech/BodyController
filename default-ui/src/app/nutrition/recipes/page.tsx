import React from "react";
import { Recipe, RecipeCard } from "./recipe_card";
import styles from './styles.module.css'


export default function Page() {
  const ingredients: Recipe[] = [
    {

      title: "Ham",
      macros: {
        calories: 100,
        proteins: 22,
        carbs: 5,
        fats: 3,
      },
    },
    {

      title: "White bread",
      macros: {
        calories: 360,
        proteins: 5.5,
        carbs: 62.4,
        fats: 6.9,
      },
    },
    {

      title: "White bread",
      macros: {
        calories: 360,
        proteins: 5.5,
        carbs: 62.4,
        fats: 6.9,
      },
    },
    {

      title: "White bread",
      macros: {
        calories: 360,
        proteins: 5.5,
        carbs: 62.4,
        fats: 6.9,
      },
    },
  ]

  return (
    <main>
      <p className={styles.page_title}>Available Ingredients</p>
      <div className={styles.grid_container}>
        {ingredients.map((value: Recipe) => {
          return (
            <RecipeCard ingredient={value} />
          )
        })}
      </div>
    </main>
  );
}
