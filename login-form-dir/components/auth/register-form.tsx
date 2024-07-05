"use client";
import * as z from "zod";
import {useForm} from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import {
Form,
FormControl,
FormField,
FormItem,
FormLabel,
FormMessage,
} from "@/components/ui/form"
import { RegisterSchema } from "@/chcemas";
import {Input} from "@/components/ui/input"
import { CardWrapper } from "@/components/auth/card-wrapper"
import { Button } from "../ui/button";

export const RegisterForm = () => {
    const form = useForm<z.infer<typeof RegisterSchema>>({
        resolver: zodResolver(RegisterSchema),
        defaultValues: {
            name: "",
            email: "",
            password: "",
        },
    });
const onSubmit = (values: z.infer<typeof RegisterSchema>) => {
    console.log(values);
    
}

    return (
        <CardWrapper
            headerLabel="Time to appear"
            backButtonLabel="Already registered?"
            backButtonHref="/auth/login"     >
            <Form {...form}>
                <form 
                onSubmit={form.handleSubmit(()=> {})}
                className="space-y-8"
                >

                    <div className="space-y-5">
<FormField
                            control={form.control}
                            name="name"
                            render = {({field}) => (
                                <FormItem>
                                    <FormLabel>
                                        Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                        {...field}
                                        placeholder=""
                                        type="name"
                                        
                                        />


                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                            />

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
                                        Pasword
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                        {...field}
                                        placeholder="******"
                                        type="password"
                                        
                                        />
                            
                            
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </div>
                    <Button
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