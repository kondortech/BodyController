import React from "react";
import { Ingredient, Tile } from "./ingredient_card";
import styles from './styles.module.css'


export default function Page() {
  const ingredients: Ingredient[] = [
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
  ]

  return (
    <main className="flex min-h-screen flex-col items-center p-12">
      {ingredients.map((value: Ingredient) => {
        return (
          <div className={styles.list_card}><Tile ingredient={value} /></div>
        )
      })}
    </main>
  );
}
