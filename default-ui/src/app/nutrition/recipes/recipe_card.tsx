"use client";

import React from 'react';
import styles from './styles.module.css'

import Image from 'next/image'

export type Macros = {
    calories: number;
    proteins: number;
    carbs: number;
    fats: number;
};

export type Recipe = {
    title: string;
    macros: Macros;
};

type Props = {
    ingredient: Recipe;

};

export const RecipeCard = (props: Props): JSX.Element => {
    return (
        <div className={styles.card}>
            <Image
                src="/ingredient-default.jpg"
                width={100}
                height={100}
                alt={props.ingredient.title}
                className={styles.card_image}
            />
            <div>
                <p className={styles.card_title}>{props.ingredient.title}</p>
                <p>Calories: {props.ingredient.macros.calories}</p>
                <p>Proteins: {props.ingredient.macros.proteins}</p>
                <p>Carbs: {props.ingredient.macros.carbs}</p>
                <p>Fats: {props.ingredient.macros.fats}</p>
            </div>
        </div >
    );
}