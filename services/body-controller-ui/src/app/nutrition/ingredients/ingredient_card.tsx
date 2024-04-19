import React from 'react';
import styles from './styles.module.css'

import Image from 'next/image'

export type Macros = {
    calories: number;
    proteins: number;
    carbs: number;
    fats: number;
};

export type Ingredient = {
    title: string;
    macros: Macros;
};

type Props = {
    ingredient: Ingredient;

};

export const Tile = (props: Props): JSX.Element => {
    return (
        <div className={styles.card}>
            <Image
                src="/default_squared.svg"
                width={100}
                height={100}
                alt="default"
                className={styles.card_image}
            />
            <div className={styles.card_info}>
                <p className={styles.title}>{props.ingredient.title}</p>
                <p>Calories: {props.ingredient.macros.calories}</p>
                <p>Proteins: {props.ingredient.macros.proteins}</p>
                <p>Carbs: {props.ingredient.macros.carbs}</p>
                <p>Fats: {props.ingredient.macros.fats}</p>
            </div>
        </div >
    );
}