"use client";
import * as z from "zod";
import { startTransition, useState, useTransition } from "react";
import {useForm} from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
Form,
FormControl,
FormField,
FormItem,
FormLabel,
FormMessage,
} from "@/components/ui/form";
import { LoginSchema } from "@/schemas";
import {Input} from "@/components/ui/input";
import { CardWrapper } from "@/components/auth/card-wrapper";
import { Button } from "../ui/button";
import { FormError } from "@/components/form-error";
import { FormSuccess } from "@/components/form-success";
import { login } from "@/actions/login";

export const LoginForm = () => {
    const [error, setError] = useState<string | undefined>("");
    const [success, setSuccess] = useState<string | undefined>("");
    const [isPending, startTransition] = useTransition();

    const form = useForm<z.infer<typeof LoginSchema>>({
        resolver: zodResolver(LoginSchema),
        defaultValues: {
            email: "",
            password: "",
        },
    });
const onSubmit = (values: z.infer<typeof LoginSchema>) => {
    setError("");
    setSuccess("");

    startTransition(() => {
        login(values)
        .then((data) => {
            setError(data.error);
            setSuccess(data.succes);
        })
    })
};

    return (
        <CardWrapper
            headerLabel="Welcome"
            backButtonLabel="Register"
            backButtonHref="/auth/register"     >
            <Form {...form}>
                <form 
                onSubmit={form.handleSubmit(onSubmit)}
                className="space-y-8"    
                >

                    <div className="space-y-5">
                        <FormField 
                            control={form.control}
                            name="email"
                            render = {({field}) => (
                                <FormItem>
                                    <FormLabel>
                                        Email
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                        {...field}
                                        disabled={isPending}
                                        placeholder="sergio.ramos@mail.ru"
                                        type="email"
                                        
                                        />
                            
                            
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />

<FormField 
                            control={form.control}
                            name="password"
                            render = {({field}) => (
                                <FormItem>
                                    <FormLabel>
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                        {...field}
                                        disabled={isPending}
                                        placeholder="******"
                                        type="password"
                                        
                                        />
                            
                            
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </div>
                    <FormError message={error}/>
                    <FormSuccess message={success}/>
                    <Button
                    disabled={isPending}
                    type = "submit"
                    className="w-full"
                    >
                        Login
                    </Button>

                </form>

            </Form>

        </CardWrapper>
    )
}