"use client";

import React from 'react';

import { ModelsIngredient } from '@/generated/services/nutrition/api';
import { deleteIngredient } from '@/services/nutrition/ingredients_api';

const defaultImagePath = "/ingredient-default.jpg";

interface IngredientCardProps {
    ingredient: ModelsIngredient;
    onClick?: (ingredient: ModelsIngredient) => void;
}

const IngredientCard: React.FC<IngredientCardProps> = ({ ingredient, onClick }) => {
    return (
        <div
            className="bg-white rounded-lg shadow-md p-4 cursor-pointer transform hover:scale-105 transition-transform"
            onClick={() => {
                if (onClick !== undefined) { // excessive -> delete this bs
                    return onClick(ingredient);
                }
            }}
        >
            <img src={defaultImagePath} alt={ingredient.title} className="w-full h-32 object-cover rounded-md mb-4" />
            <h2 className="text-lg font-bold mb-2">{ingredient.title}</h2>
            <p>Calories: {ingredient.macrosNormalized.calories}</p>
            <p>Proteins: {ingredient.macrosNormalized.proteins}</p>
            <p>Carbs: {ingredient.macrosNormalized.carbs}</p>
            <p>Fat: {ingredient.macrosNormalized.fats}</p>
            <div onClick={() => {
                deleteIngredient(ingredient.id);
            }} className="w-full h-12 object-cover rounded-md mb-4 bg-red-600">
                delete
            </div>
        </div>
    );
};

export default IngredientCard;