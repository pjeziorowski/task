import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    let userId = event.cookies.get("auth");
    if (!userId) {
        userId = crypto.randomUUID();
        event.cookies.set("auth", userId, {
            path: "/",
            httpOnly: true,
            sameSite: "strict",
            maxAge: 60 * 60 * 24 * 365,
        });
    }
    
    event.locals.userId = userId;

    return resolve(event);
};
