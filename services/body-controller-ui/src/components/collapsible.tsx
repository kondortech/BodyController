'use client';
import React, { ReactNode } from 'react';

import { useEffect, useState } from "react";

type props = {
    title: string;
    children: ReactNode;
};

export const Collapsible = (props: props) => {
    const [open, setOpen] = useState(false);

    const onToggle = () => {
        setOpen((prev) => !prev);
    }

    return (
        <>
            <button onClick={onToggle}>{props.title}</button>
            {open && props.children}
        </>
    )
}

