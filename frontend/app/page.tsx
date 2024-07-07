import { LoginButton } from "@/components/auth/login-button";
import {Button} from "@/components/ui/button";
export default function Home() {
  return ( 
<main className="flex h-full flex-col items-center justify-center">
  <div className="space-y-6 text-center">
    <h1 className="text-3xl font semibold ">
      Ramos Project
    </h1>
    <p>
      planning platform
    </p>

  </div>
  <div>
    <LoginButton>
    <Button variant= "secondary" size = "lg">
      Sign in
    </Button>
    </LoginButton>
  </div>
</main>

)
}
