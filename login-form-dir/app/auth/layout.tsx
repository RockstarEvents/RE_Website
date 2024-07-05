const AuthLayout =({
    children
} : {
    children: React.ReactNode
}
) => {
    return (
        <div className="h-full flex items-center justify-center bg-green-400">
            {children}
        </div>

    );
}
export default AuthLayout;