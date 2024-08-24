"use client";

import React, { useState } from 'react';

import { ModelsIngredient } from '@/generated/services/nutrition/api';

const defaultImagePath = "/ingredient-default.jpg";

interface IngredientCardProps {
    ingredient: ModelsIngredient;
    onSelect: (ingredient: ModelsIngredient) => void;
    onDeselect: (ingredient: ModelsIngredient) => void;
}

const IngredientCard: React.FC<IngredientCardProps> = ({ ingredient, onSelect, onDeselect }) => {
    const [selected, setSelected] = useState<boolean>(false);

    return (
        <div
            className={"bg-white rounded-lg shadow-md p-4 cursor-pointer transform hover:scale-105 transition-transform border" + (selected ? " border-green-500" : "")}
            onClick={() => {
                if (!selected) {
                    onSelect(ingredient);
                } else {
                    onDeselect(ingredient);
                }
                setSelected((prevValue: boolean) => !prevValue);
            }}
        >
            <img src={defaultImagePath} alt={ingredient.title} className="w-full h-32 object-cover rounded-md mb-4" />
            <h2 className="text-lg font-bold mb-2">{ingredient.title}</h2>
            <p>Calories: {ingredient.macrosNormalized.calories}</p>
            <p>Proteins: {ingredient.macrosNormalized.proteins}</p>
            <p>Carbs: {ingredient.macrosNormalized.carbs}</p>
            <p>Fat: {ingredient.macrosNormalized.fats}</p>
        </div>
    );
};

export default IngredientCard;