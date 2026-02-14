import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { FileQuestion, Home, MoveLeft } from "lucide-react";

const NotFound = () => {
  return (
    <div className="min-h-screen w-full flex items-center justify-center bg-slate-50 p-4">
      <Card className="w-full max-w-md shadow-lg border-slate-200">
        <CardHeader className="text-center pb-2">
          <div className="flex justify-center mb-4">
            <div className="p-3 bg-slate-100 rounded-full">
              <FileQuestion className="h-10 w-10 text-slate-600" />
            </div>
          </div>
          <CardTitle className="text-7xl font-bold tracking-tighter text-slate-900">
            404
          </CardTitle>
          <p className="text-xl font-medium text-slate-600 mt-2">
            Page Not Found
          </p>
        </CardHeader>

        <CardContent className="text-center py-6">
          <p className="text-slate-500 leading-relaxed">
            Oops! The page you're looking for doesn't exist or has been moved to
            a different URL.
          </p>
        </CardContent>

        <CardFooter className="flex flex-col sm:flex-row gap-3 justify-center pt-2">
          <Button asChild variant="outline" className="w-full sm:w-auto">
            <Link to={-1 as any} className="flex items-center gap-2">
              <MoveLeft className="h-4 w-4" />
              Go Back
            </Link>
          </Button>
          <Button asChild className="w-full sm:w-auto">
            <Link to="/" className="flex items-center gap-2">
              <Home className="h-4 w-4" />
              Back to Home
            </Link>
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
};

export default NotFound;
