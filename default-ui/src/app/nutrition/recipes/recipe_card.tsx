"use client";

import React from 'react';

const defaultImagePath = "/ingredient-default.jpg";

export type Props = {
    title: string;
    description: string;
};

export const RecipeCard = (props: Props): JSX.Element => {
    return (
        <div className="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl">
            <div className="md:flex">
                <div className="md:shrink-0">
                    <img className="h-48 w-full object-cover md:h-full md:w-48" src={defaultImagePath} alt={props.title} />
                </div>
                <div className="p-8">
                    <div className="uppercase tracking-wide text-sm text-indigo-500 font-semibold">{props.title}</div>
                    <p className="mt-2 text-gray-500">{props.description}</p>
                </div>
            </div>
        </div>
    );
}