"use client";

import { ModelsRecipe } from '@/generated/services/nutrition/api';
import { deleteRecipe } from '@/services/nutrition/api';
import React from 'react';

const defaultImagePath = "/ingredient-default.jpg";

export type Props = {
    recipe: ModelsRecipe;
};

export const RecipeCard = (props: Props): JSX.Element => {
    return (
        <div className="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl">
            <div className="md:flex">
                <div className="md:shrink-0">
                    <img className="h-48 w-full object-cover md:h-full md:w-48" src={defaultImagePath} alt={props.recipe.title} />
                </div>
                <div className="p-8">
                    <div className="uppercase tracking-wide text-sm text-indigo-500 font-semibold">{props.recipe.title}</div>
                    <p className="mt-2 text-gray-500">{props.recipe.description}</p>
                </div>
                <button
                    onClick={() => { deleteRecipe(props.recipe.id) }}
                    className="flex justify-between border text-red-500 hover:text-red-700 focus:outline-none"
                    aria-label="Delete"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="h-6 w-6"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M6 18L18 6M6 6l12 12"
                        />
                    </svg>
                </button>
            </div>
        </div >
    );
}