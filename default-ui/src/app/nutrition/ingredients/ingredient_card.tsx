"use client";

import React from 'react';
import styles from './styles.module.css'

import Image from 'next/image'
import { ModelsIngredient } from '@/generated/services/nutrition/api';

type Props = {
    ingredient: ModelsIngredient;

};

export const IngredientCard = (props: Props): JSX.Element => {
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
                <p>Calories: {props.ingredient.macrosNormalized.calories}</p>
                <p>Proteins: {props.ingredient.macrosNormalized.proteins}</p>
                <p>Carbs: {props.ingredient.macrosNormalized.carbs}</p>
                <p>Fats: {props.ingredient.macrosNormalized.fats}</p>
            </div>
        </div >
    );
}